package data

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"os"
)

type Chunk struct {
	Path string
	Hash string
	Num  uint64
}

//go:generate mockery --name=ChunkWriter --case underscore
type ChunkWriter func(chunk *Chunk, chunkBytes *[]byte) error

func NewChunk(path string, hash string, number uint64) *Chunk {
	return &Chunk{Path: path, Hash: hash, Num: number}
}

func DefaultChunkWriter(chunk *Chunk, chunkBytes *[]byte) error {
	// write/save buffer to disk
	err := afero.WriteFile(AppFS, chunk.Path, *chunkBytes, os.ModeAppend)
	if err != nil {
		log.Errorf("Chunk writer Failed to write the chunk data to a new file: %v", err)
		return err
	}

	return nil
}
