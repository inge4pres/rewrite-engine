package rewrite_test

import (
	"bytes"
	"testing"

	. "git.bravofly.com/fgualazzi/rewrite-engine.git/pkg/rewrite"
	"github.com/stretchr/testify/assert"
)

const exampleConf = `{ 
	"rules" : [
		{
			"http_status": 302,
			"source": "/",
			"target":"http://newhost.com/context"
		},
		{
			"http_status": 301,
			"source": "/watch",
			"target":"https://watch.com/"
		}
	]
}`
const wrongConf = ` {
	"rules": [
		{
			{
				"target":"http://newhost.com/context"
			}
		}
	]
}`

var correctConfReader = bytes.NewBuffer([]byte(exampleConf))
var wrongConfReader = bytes.NewBuffer([]byte(wrongConf))

func TestParseCorrectConfig(t *testing.T) {
	conf, err := ParseJSONConfig(correctConfReader)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(conf.Rules))
}

func TestParseWrongConfig(t *testing.T) {
	_, err := ParseJSONConfig(wrongConfReader)
	assert.NotNil(t, err)
}
