//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package replication

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/weaviate/weaviate/entities/models"
	"github.com/weaviate/weaviate/test/docker"
	"github.com/weaviate/weaviate/test/helper"
	"github.com/weaviate/weaviate/test/helper/sample-schema/articles"
	"github.com/weaviate/weaviate/usecases/replica"
)

func asyncRepairSimpleScenario(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	compose, err := docker.New().
		WithWeaviateCluster(3).
		WithText2VecContextionary().
		Start(ctx)
	require.Nil(t, err)
	defer func() {
		if err := compose.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate test containers: %s", err.Error())
		}
	}()

	helper.SetupClient(compose.GetWeaviate().URI())
	paragraphClass := articles.ParagraphsClass()
	articleClass := articles.ArticlesClass()

	t.Run("create schema", func(t *testing.T) {
		paragraphClass.ReplicationConfig = &models.ReplicationConfig{
			Factor:       3,
			AsyncEnabled: true,
		}
		paragraphClass.Vectorizer = "text2vec-contextionary"
		helper.CreateClass(t, paragraphClass)
		articleClass.ReplicationConfig = &models.ReplicationConfig{
			Factor:       3,
			AsyncEnabled: true,
		}
		helper.CreateClass(t, articleClass)
	})

	t.Run("insert paragraphs", func(t *testing.T) {
		batch := make([]*models.Object, len(paragraphIDs))
		for i, id := range paragraphIDs {
			batch[i] = articles.NewParagraph().
				WithID(id).
				WithContents(fmt.Sprintf("paragraph#%d", i)).
				Object()
		}
		createObjects(t, compose.GetWeaviate().URI(), batch)
	})

	t.Run("insert articles", func(t *testing.T) {
		batch := make([]*models.Object, len(articleIDs))
		for i, id := range articleIDs {
			batch[i] = articles.NewArticle().
				WithID(id).
				WithTitle(fmt.Sprintf("Article#%d", i)).
				Object()
		}
		createObjects(t, compose.GetWeaviateNode(2).URI(), batch)
	})

	t.Run("stop node 3", func(t *testing.T) {
		stopNodeAt(ctx, t, compose, 3)
	})

	repairObj := models.Object{
		ID:    "e5390693-5a22-44b8-997d-2a213aaf5884",
		Class: "Paragraph",
		Properties: map[string]interface{}{
			"contents": "a new paragraph",
		},
	}

	t.Run("add new object to node one", func(t *testing.T) {
		createObjectCL(t, compose.GetWeaviate().URI(), &repairObj, replica.One)
	})

	t.Run("restart node 3", func(t *testing.T) {
		startNodeAt(ctx, t, compose, 3)
		time.Sleep(time.Second)
	})

	// wait for some time for async replication to repair missing object
	time.Sleep(3 * time.Second)

	t.Run("stop node 2", func(t *testing.T) {
		stopNodeAt(ctx, t, compose, 2)
	})

	t.Run("assert new object read repair was made", func(t *testing.T) {
		resp, err := getObjectCL(t, compose.GetWeaviateNode(3).URI(),
			repairObj.Class, repairObj.ID, replica.One)
		require.Nil(t, err)
		assert.Equal(t, repairObj.ID, resp.ID)
		assert.Equal(t, repairObj.Class, resp.Class)
		assert.EqualValues(t, repairObj.Properties, resp.Properties)
		assert.EqualValues(t, repairObj.Vector, resp.Vector)
	})

	replaceObj := repairObj
	replaceObj.Properties = map[string]interface{}{
		"contents": "this paragraph was replaced",
	}

	t.Run("replace object", func(t *testing.T) {
		updateObjectCL(t, compose.GetWeaviateNode(3).URI(), &replaceObj, replica.One)
	})

	t.Run("restart node 2", func(t *testing.T) {
		startNodeAt(ctx, t, compose, 2)
		time.Sleep(time.Second)
	})

	// wait for some time for async replication to repair missing object
	time.Sleep(3 * time.Second)

	t.Run("stop node 3", func(t *testing.T) {
		stopNodeAt(ctx, t, compose, 3)
	})

	t.Run("assert updated object read repair was made", func(t *testing.T) {
		exists, err := objectExistsCL(t, compose.GetWeaviateNode(2).URI(),
			replaceObj.Class, replaceObj.ID, replica.One)
		require.Nil(t, err)
		require.True(t, exists)

		resp, err := getObjectCL(t, compose.GetWeaviate().URI(),
			repairObj.Class, repairObj.ID, replica.One)
		require.Nil(t, err)
		assert.Equal(t, replaceObj.ID, resp.ID)
		assert.Equal(t, replaceObj.Class, resp.Class)
		assert.EqualValues(t, replaceObj.Properties, resp.Properties)
		assert.EqualValues(t, replaceObj.Vector, resp.Vector)
	})
}
