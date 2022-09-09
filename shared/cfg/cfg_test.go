package cfg

import "testing"

func TestNew(t *testing.T) {
	var c, d *ConfigValue

	for i := 0; i < 100; i++ {
		c, d = new(ConfigValue), new(ConfigValue)
		c.New()
		d.New()

		if *c != *d {
			t.Error("*ConfigValue.New() does not present the same behaviour.")
		}

		c, d = nil, nil
	}
}

func TestUseSQLite(t *testing.T) {
	c := new(ConfigValue)
	c.New()

	if !c.UseSQLite() {
		t.Errorf("Expected 'true'; Got 'false\n")
	}
}
