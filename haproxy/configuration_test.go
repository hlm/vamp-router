package haproxy

import (
	"os"
	"testing"
)

const (
	TEMPLATE_FILE = "../configuration/templates/haproxy_config.template"
	CONFIG_FILE   = "/tmp/vamp_lb_test.cfg"
	EXAMPLE       = "../test/test_config1.json"
	JSON_FILE     = "/tmp/vamp_lb_test.json"
	PID_FILE      = "/tmp/vamp_lb_test.pid"
)

var (
	haConfig = Config{TemplateFile: TEMPLATE_FILE, ConfigFile: CONFIG_FILE, JsonFile: JSON_FILE, PidFile: PID_FILE}
)

func TestConfiguration_GetConfigFromDisk(t *testing.T) {
	err := haConfig.GetConfigFromDisk(EXAMPLE)
	if err != nil {
		t.Errorf("err: %v", err)
	}

	err = haConfig.GetConfigFromDisk("/this_is_really_something_wrong")
	if err == nil {
		t.Errorf("Expected an error")
	}
}

func TestConfiguration_SetWeight(t *testing.T) {
	err := haConfig.SetWeight("test_be_1", "test_be_1_a", 20)
	if err != nil {
		t.Errorf("err: %v", err)
	}
}

func TestConfiguration_GetFrontends(t *testing.T) {
	result := haConfig.GetFrontends()
	if result[0].Name != "test_fe_1" {
		t.Errorf("Failed to get frontends array")
	}

}

func TestConfiguration_GetFrontend(t *testing.T) {
	result := haConfig.GetFrontend("test_fe_1")
	if result.Name != "test_fe_1" {
		t.Errorf("Failed to get frontend")
	}

}

func TestConfiguration_AddFrontend(t *testing.T) {

	fe := Frontend{Name: "my_test_frontend", Mode: "http", DefaultBackend: "test_be_1"}
	err := haConfig.AddFrontend(&fe)
	if err != nil {
		t.Errorf("Failed to add frontend")
	}
	if haConfig.Frontends[3].Name != "my_test_frontend" {
		t.Errorf("Failed to add frontend")
	}
}

func TestConfiguration_DeleteFrontend(t *testing.T) {

	result := haConfig.DeleteFrontend("non_existing_backend")
	if result != false {
		t.Errorf("Backend should not be removed")
	}

	result = haConfig.DeleteFrontend("test_fe_2")
	if result != true {
		t.Errorf("Failed to remove frontend")
	}

}

func TestConfiguration_GetAcls(t *testing.T) {

	acls := haConfig.GetAcls("test_fe_1")
	if acls[0].Name != "uses_internetexplorer" {
		t.Errorf("Could not retrieve ACL")
	}
}

func TestConfiguration_AddAcl(t *testing.T) {

	acl := ACL{Name: "uses_firefox", Backend: "test_be_1_b", Pattern: "hdr_sub(user-agent) Mozilla"}
	err := haConfig.AddAcl("test_fe_1", &acl)
	if err != nil {
		t.Errorf("Could not add ACL")
	}
	if haConfig.Frontends[0].ACLs[1].Name != "uses_firefox" {
		t.Errorf("Could not add ACL")
	}
}

func TestConfiguration_Render(t *testing.T) {
	err := haConfig.Render()
	if err != nil {
		t.Errorf("err: %v", err)
	}
}

func TestConfiguration_Persist(t *testing.T) {
	err := haConfig.Persist()
	if err != nil {
		t.Errorf("err: %v", err)
	}
	os.Remove(CONFIG_FILE)
	os.Remove(JSON_FILE)
}

func TestConfiguration_RenderAndPersist(t *testing.T) {
	err := haConfig.RenderAndPersist()
	if err != nil {
		t.Errorf("err: %v", err)
	}
	os.Remove(CONFIG_FILE)
	os.Remove(JSON_FILE)
}
