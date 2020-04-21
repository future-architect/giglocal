package gcplocal_test

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"time"

	"cloud.google.com/go/datastore"
)

type User struct {
	Id  string
	Age int
}

var (
	dbPath = filepath.FromSlash("../src/datastore/.data/WEB-INF/appengine-generated/local_db.bin")
	ctx    = context.Background()
	dc     *datastore.Client
)

func init() {
	os.Setenv("DATASTORE_EMULATOR_HOST", DataStoreEmulatorHost)
}

func TestDatastore(t *testing.T) {
	// テスト用データ //
	const userID1 = "1"
	const userID2 = "2"
	var inUser1 = User{Id: "hoge", Age: 22}
	var inUser2 = User{Id: "fuga", Age: 24}

	////// 0. setting //////
	// dc の作成
	var err error
	dc, err = datastore.NewClient(ctx, projectID)
	if err != nil {
		t.Fatal(err)
	}

	////// TEST1. データ作成 //////
	// userID={1|2} <= inUser{1|2} を追加
	t.Run("TEST1", func(t *testing.T) {
		if err := Create(inUser1, userID1); err != nil {
			t.Fatalf("[ERROR] Failed to save task: %v", err)
		}
		if err := Create(inUser2, userID2); err != nil {
			t.Fatalf("[ERROR] Failed to save task: %v", err)
		}
	})

	////// TEST2. データの取得，整合性の確認 //////
	// userID={1|2} => gotUser{1|2} を取得
	// inUser{1|2} == gotUser{1|2} であることを確認
	t.Run("TEST2", func(t *testing.T) {
		gotUser1, err := Get(userID1)
		if err != nil {
			t.Fatalf("[ERROR] Failed to get task: %v", err)
		}
		gotUser2, err := Get(userID2)
		if err != nil {
			t.Fatalf("[ERROR] Failed to get task: %v", err)
		}

		want2 := []User{inUser1, inUser2}
		real2 := []User{gotUser1, gotUser2}
		if !reflect.DeepEqual(want2, real2) {
			t.Logf("[ERROR] want %v", want2)
			t.Logf("[ERROR] real %v", real2)
			t.Fatalf("[ERROR] NOT match data")
		}
	})

	////// TEST3. データの削除，整合性の確認 //////
	// userID=1 => 削除
	// GetAll() => gotUsers を取得
	// gotUsers == {inUser2} であることを確認
	t.Run("TEST3", func(t *testing.T) {
		if err := Delete(userID1); err != nil {
			t.Fatalf("[ERROR] Failed to delete task: %v", err)
		}

		want3 := []User{inUser2}
		real3, err := GetAll()
		if err != nil {
			t.Fatalf("[ERROR] Failed to save task: %v", err)
		}

		if !reflect.DeepEqual(want3, real3) {
			t.Logf("[ERROR] want %v", want3)
			t.Logf("[ERROR] real %v", real3)
			t.Fatalf("[ERROR] NOT match data")
		}
	})

	// local_db.bin が生成されるまでに時間がかかるので待つ
	for {
		if _, err := os.Stat(dbPath); err == nil {
			break
		}
		time.Sleep(10 * time.Second)
	}
	// db にデータが反映されるように適当に余裕を持たせる
	time.Sleep(30 * time.Second)

	//// TEST4. エミュの再起動，整合性の確認 //////
	// コンテナ再起動
	// gotUsers == {inUser2} であることを確認
	// (ちゃんとマウントされていれば永続化されるため一致する)
	t.Run("TEST4", func(t *testing.T) {
		if err := RestartContainer(); err != nil {
			t.Fatalf("[ERROR] Failed to docker-compose stop: %v", err)
		}

		want4 := []User{inUser2}
		real4, err := GetAll()
		if err != nil {
			t.Fatalf("[ERROR] Failed to save task: %v", err)
		}
		if !reflect.DeepEqual(want4, real4) {
			t.Logf("[ERROR] want %v", want4)
			t.Logf("[ERROR] real %v", real4)
			t.Fatalf("[ERROR] NOT match data")
		}
	})
}

func Create(user User, userID string) error {
	_, err := dc.Put(ctx, datastore.NameKey("User", userID, nil), &user)
	return err
}

func Get(userID string) (User, error) {
	user := User{}
	err := dc.Get(ctx, datastore.NameKey("User", userID, nil), &user)
	return user, err
}

func Delete(userID string) error {
	return dc.Delete(ctx, datastore.NameKey("User", userID, nil))
}

func GetAll() ([]User, error) {
	var users []User
	_, err := dc.GetAll(ctx, datastore.NewQuery("User"), &users)
	return users, err
}

func RestartContainer() error {
	cmd := exec.Command("docker-compose", "restart")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "cd", "..", "&", "docker-compose", "restart")
	}
	return cmd.Run()
}
