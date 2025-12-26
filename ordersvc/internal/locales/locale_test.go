package locales

import (
	"context"
	"testing"
)

func TestGetLocalize(t *testing.T) {
	localizer, err := GetLocalize(context.Background(), "zh", OrderCreateSuccess)
	if err != nil {
		t.Errorf("GetLocalize() error = %v", err)
		return
	}
	if localizer == "" {
		t.Errorf("GetLocalize() localizer = %v", localizer)
		return
	}
}

func TestGetLocaleMsg(t *testing.T) {
	msg := GetLocaleMsg(context.Background(), "zh", OrderCreateSuccess)
	if msg == "" {
		t.Errorf("GetLocaleMsg() msg = %v", msg)
		return
	}
	t.Logf("msg: %s", msg)
}
