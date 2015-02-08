package quickunionweighted

import(
  "github.com/stretchr/testify/assert"
  "github.com/iloire/quickunionweighted"
  "testing"
)

func TestUnion(t *testing.T) { 
  q:= new(quickunionweighted.QuickUnion)
  q.Initialise(10)

  assert.False(t, q.IsConnected(8,0))
  q.Union(8, 0)
  assert.True(t, q.IsConnected(8,0))

  assert.False(t, q.IsConnected(1,2))
  q.Union(1, 2)
  assert.True(t, q.IsConnected(1,2))

  q.Union(1, 8)
  assert.True(t, q.IsConnected(8,2))
}