package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOk(t *testing.T) {
	// Создаем новый HTTP-запрос
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=2", nil)
	// Создаем новый записывающий объект для ответа
	rr := httptest.NewRecorder()
	// Вызываем обработчик
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(rr, req)
	// Проверяем код ответа
	require.Equal(t, http.StatusOK, rr.Code)

	// Проверяем, что тело ответа не пустое
	expectedLength := len(rr.Body.String())
	assert.Greater(t, expectedLength, 0)
}

func TestMainHandlerWhenUnsupportedCity(t *testing.T) {
	// Создаем новый HTTP-запрос
	req := httptest.NewRequest("GET", "/cafe?city=unknown&count=2", nil)
	// Создаем новый записывающий объект для ответа
	rr := httptest.NewRecorder()
	// Вызываем обработчик
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(rr, req)
	// Проверяем код ответа
	require.Equal(t, http.StatusBadRequest, rr.Code)
	// Проверяем содержимое ответа
	assert.Equal(t, "wrong city value", rr.Body.String())
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	// Создаем новый HTTP-запрос
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=10", nil)
	// Создаем новый записывающий объект для ответа
	rr := httptest.NewRecorder()
	// Вызываем обработчик
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(rr, req)
	// Проверяем код ответа
	require.Equal(t, http.StatusOK, rr.Code)

	// Проверяем, что количество кафе в ответе соответствует totalCount
	actualCafes := strings.Split(rr.Body.String(), ",")
	assert.Equal(t, totalCount, len(actualCafes), "Количество кафе в ответе не соответствует ожидаемому")

	// Проверяем, что ответ содержит все доступные кафе
	expectedBody := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"
	assert.Equal(t, expectedBody, rr.Body.String())
}
