package gcplocal_test

import (
	"context"
	"os"
	"testing"

	"cloud.google.com/go/firestore"
)

const (
	CollectionDocument = "collection/document"
)

type dataMap struct {
	EmulatorName string `firestore:"emi"`
	Result       string `firestore:"res"`
}

var data = dataMap{
	EmulatorName: "firestore",
	Result:       "success",
}

func init() {
	os.Setenv("FIRESTORE_EMULATOR_HOST", FireStoreEmulatorHost)
}

func TestFirestore(t *testing.T) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	// Write
	if _, err = client.Doc(CollectionDocument).Create(ctx, data); err != nil {
		t.Fatal(err)
	}

	// Read
	var nyData dataMap
	docsnap, err := client.Doc(CollectionDocument).Get(ctx)
	if err := docsnap.DataTo(&nyData); err != nil {
		t.Fatal(err)
	}

	// Verify read data
	if data != nyData {
		t.Fatal("data didn't match")
	}

	// Delete write data
	if _, err := client.Doc(CollectionDocument).Delete(ctx); err != nil {
		t.Fatal(err)
	}
}
