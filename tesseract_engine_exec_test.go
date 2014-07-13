package ocrworker

import (
	"encoding/json"
	"testing"

	"io/ioutil"

	"github.com/couchbaselabs/go.assert"
	"github.com/couchbaselabs/logg"
)

func TestTesseractEngineExecWithRequest(t *testing.T) {

	engine := TesseractEngineExec{}
	bytes, err := ioutil.ReadFile("docs/testimage.png")
	assert.True(t, err == nil)

	cFlags := make(map[string]interface{})
	cFlags["tessedit_char_whitelist"] = "0123456789"

	ocrRequest := OcrRequest{
		ImgBytes:   bytes,
		EngineType: ENGINE_TESSERACT_EXEC,
		EngineArgs: cFlags,
	}

	assert.True(t, err == nil)
	result, err := engine.ProcessRequest(ocrRequest)
	assert.True(t, err == nil)
	logg.LogTo("TEST", "result: %v", result)

}

func TestTesseractEngineExecWithJson(t *testing.T) {

	testJson := `{"engine":"tesseract_exec", "engine_args":{"c_flags":{"tessedit_char_whitelist":"0123456789"}}}`
	ocrRequest := OcrRequest{}
	err := json.Unmarshal([]byte(testJson), &ocrRequest)
	assert.True(t, err == nil)
	bytes, err := ioutil.ReadFile("docs/testimage.png")
	assert.True(t, err == nil)
	ocrRequest.ImgBytes = bytes
	engine := NewOcrEngine(ocrRequest.EngineType)
	result, err := engine.ProcessRequest(ocrRequest)
	logg.LogTo("TEST", "err: %v", err)
	assert.True(t, err == nil)
	logg.LogTo("TEST", "result: %v", result)

}

func TestTesseractEngineExecWithFile(t *testing.T) {

	engine := TesseractEngineExec{}
	engineArgs := TesseractEngineExecArgs{}
	result, err := engine.processImageFile("docs/testimage.png", engineArgs)
	assert.True(t, err == nil)
	logg.LogTo("TEST", "result: %v", result)

}
