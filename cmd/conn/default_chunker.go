package conn

import "github.com/kris-nova/rtmp/chunk"

// DefaultChunker provides a default implementation of the Chunker interface.
// All chunks generated are sent over the given StreamId.
type DefaultChunker struct {
	// StreamId is the uint32 chunk stream ID that will be attached to the
	// basic header of each chunk generated by the Chunk method.
	StreamId uint32
}

// NewChunker returns a new instance of the Chunker interface, using the
// DefaultChunker as an implementation. It initializes the DefaultChunker with
// the given stream id, and returns it.
func NewChunker(sid uint32) Chunker {
	return &DefaultChunker{
		StreamId: sid,
	}
}

// Chunk implements the Chunk method on Chunker.Chunk. It marshals the data in
// the given ConnSendable, using the amf0/encoding package, and writes it into a
// chunk to be sent over the correct chunk stream ID.
func (c *DefaultChunker) Chunk(m Marshallable) (*chunk.Chunk, error) {
	data, err := m.Marshal()
	if err != nil {
		return nil, err
	}

	return &chunk.Chunk{
		Header: &chunk.Header{
			BasicHeader: chunk.BasicHeader{
				FormatId: 0,
				StreamId: ChunkStreamId,
			},
			MessageHeader: chunk.MessageHeader{
				TypeId: 0x14,
				Length: uint32(len(data)),
			},
		},
		Data: data,
	}, nil
}
