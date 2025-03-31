package htmlintegrity

import "testing"

func TestIntegirtyString256(t *testing.T) {
	result := IntegirtyString([]byte("test"), 256)
	expected := "sha256-n4bQgYhMfWWaL+qgxVrQFaO/TxsrC4Is0V1sFbDwCgg="
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestIntegirtyString384(t *testing.T) {
	result := IntegirtyString([]byte("test"), 384)
	expected := "sha384-doQSMg97CqWBL85CjcRwazyuUOAqZMqhangiSb/o78S37xzLEmJV0ZYEff7fF6Cp"
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}
