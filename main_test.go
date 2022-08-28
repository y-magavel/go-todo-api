package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"net"
	"net/http"
	"testing"
)

func TestRun(t *testing.T) {
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return run(ctx, l)
	})

	in := "message"
	url := fmt.Sprintf("http://%s/%s", l.Addr().String(), in)

	// どんなポート番号でリッスンしているのか確認
	t.Logf("try request to %q", url)

	rsp, err := http.Get(url)
	if err != nil {
		t.Errorf("faild to get: %+v", err)
	}
	defer rsp.Body.Close()
	got, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		t.Errorf("faild to read body: %v", err)
	}

	// HTTPサーバーの戻り値を検証する
	want := fmt.Sprintf("Hello, %s!", in)
	if string(got) != want {
		t.Errorf("want %s, but got %s", want, got)
	}

	// run関数に終了通知を送信する。
	cancel()

	// run関数の戻り値を検証する
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}
