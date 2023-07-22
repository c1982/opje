package locator

import "testing"

type TestServiceInterface interface {
	Call() string
}

type TestServiceInterface2 interface {
	Me() string
}

type TestService2 struct {
}

func (t *TestService2) Me() string {
	return "test2"
}

type TestService struct {
}

func (t *TestService) Call() string {
	return "test"
}

func newTestService2() TestServiceInterface2 {
	return &TestService2{}
}

func newTestService1() TestServiceInterface {
	return &TestService{}
}

func TestGetTypeName(t *testing.T) {
	l := newServiceLocator()
	name := l.getTypeName(newTestService1())
	if name != "*locator.TestService" {
		t.Errorf("expected locator.TestService, got %s", name)
	}

	t.Run("TestNameUniqueness", func(t *testing.T) {
		s1 := struct {
			Name string
			Age  int
		}{
			Name: "",
			Age:  0,
		}

		s2 := struct {
			Name2 string
			Fake  byte
		}{
			Name2: "",
			Fake:  0,
		}

		name1 := l.getTypeName(s1)
		name2 := l.getTypeName(s2)

		if name1 == name2 {
			t.Errorf("name not working, got %s and %s", name1, name2)
		}
	})
}

func TestRegisterAndResolve(t *testing.T) {
	Register(newTestService1())
	Register(newTestService2())

	srv, err := Resolve[TestServiceInterface]()
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	if srv.Call() != "test" {
		t.Errorf("expected test, got %s", srv.Call())
	}
}

func TestListCount(t *testing.T) {
	Register(newTestService1())
	Register(newTestService1())
	Register(newTestService2())

	services := List()

	if len(services) != 2 {
		t.Errorf("expected 2 services, got %d", len(services))
	}
}
