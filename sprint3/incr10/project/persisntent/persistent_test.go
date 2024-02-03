// файл persistent/persistent_test.go

package persistent

import (
	"errors"
	"testing"

	mock_store "github.com/dmad1989/ypracticumTasks/sprint3/incr10/project/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	// создаём контроллер
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// создаём объект-заглушку
	m := mock_store.NewMockStore(ctrl)
	errEmptyKey := errors.New("Указан пустой ключ")
	// гарантируем, что заглушка
	// при вызове с аргументом "Key" вернёт "Value"
	// value := []byte("Value")
	// m.EXPECT().Get("Key").Return(value, nil)
	m.EXPECT().Get("").Return([]byte(""), errEmptyKey)

	// тестируем функцию Lookup, передав в неё объект-заглушку
	_, err := Lookup(m, "")
	// и проверяем возвращаемые значения
	require.ErrorIs(t, err, errEmptyKey)
}
