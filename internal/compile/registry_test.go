package compile

import "testing"

func TestQwenAliasesSharePlan(t *testing.T) {
	base, ok := Lookup(QwenPattern)
	if !ok {
		t.Fatal("P05 missing")
	}
	for _, p := range []string{QwenLiteralNewlinesPattern, QwenPossessivePattern} {
		plan, ok := Lookup(p)
		if !ok || plan != base {
			t.Fatal("validated alias must reuse the immutable P05 plan")
		}
	}
}
