package wasimalloc

import (
	"context"
	_ "embed"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

//go:embed testdata/main.wasm
var wasm []byte

func TestMalloc(t *testing.T) {
	ctx := context.Background()
	r := wazero.NewRuntimeWithConfig(ctx, wazero.NewRuntimeConfig().WithWasmCore2())
	defer r.Close(ctx)

	_, err := wasi_snapshot_preview1.Instantiate(ctx, r)
	require.NoError(t, err)

	mod, err := r.InstantiateModuleFromBinary(ctx, wasm)
	require.NoError(t, err)

	getBuf := mod.ExportedFunction("get_buf")
	releaseBuf := mod.ExportedFunction("release_buf")
	work := mod.ExportedFunction("work")

	ret, err := getBuf.Call(ctx)
	require.NoError(t, err)
	defer releaseBuf.Call(ctx, ret[0])

	bufPtr := uint32(ret[0])
	mod.Memory().Write(ctx, bufPtr, []byte("bear"))

	ret, err = work.Call(ctx, 0)
	require.NoError(t, err)

	buf, ok := mod.Memory().Read(ctx, uint32(ret[0]), 9)
	require.EqualValues(t, "pandabear", buf)

	buf, ok = mod.Memory().Read(ctx, bufPtr, 4)
	require.True(t, ok)
	require.EqualValues(t, "bear", buf)
}
