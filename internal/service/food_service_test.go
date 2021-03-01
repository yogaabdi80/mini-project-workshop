package service

import (
	"fmt"
	"testing"
	"workshop-golang-psm/food-app-service/internal/domain/entity"
	"workshop-golang-psm/food-app-service/internal/service/mocks"

	"github.com/golang/mock/gomock"
)

func TestFoodService_SaveFood(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockFood := mocks.NewMockIFoodRepository(mockCtrl)
	testFood := &FoodService{foodRepo: mockFood}

	mockFood.EXPECT().SaveFood(gomock.Any()).Return(nil, nil)
	_, err := testFood.SaveFood(&entity.FoodViewModel{})

	if err != nil {
		errorMessage := fmt.Sprintf("Expected: %s, got: %s", "nil", err.Error())
		t.Error(errorMessage)
	}
}

func TestFoodService_GetDetailFood(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockFood := mocks.NewMockIFoodRepository(mockCtrl)
	testFood := &FoodService{foodRepo: mockFood}

	mockFood.EXPECT().GetDetailFood(1).Return(nil, nil)
	_, err := testFood.GetDetailFood(1)

	if err != nil {
		errorMessage := fmt.Sprintf("Expected: %s, got: %s", "nil", err.Error())
		t.Error(errorMessage)
	}
}

func TestFoodService_GetAllFood(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockFood := mocks.NewMockIFoodRepository(mockCtrl)
	testFood := &FoodService{foodRepo: mockFood}

	mockFood.EXPECT().GetAllFood().Return(nil, nil)
	_, err := testFood.GetAllFood()
	if err != nil {
		errorMessage := fmt.Sprintf("Expected: %s, got: %s", "nil", err.Error())
		t.Error(errorMessage)
	}

}

func TestFoodService_UpdateFood(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockFood := mocks.NewMockIFoodRepository(mockCtrl)
	testFood := &FoodService{foodRepo: mockFood}

	mockFood.EXPECT().UpdateFood(gomock.Any()).Return(nil, nil)
	_, err := testFood.UpdateFood(&entity.FoodViewModel{})
	if err != nil {
		errorMessage := fmt.Sprintf("Expected: %s, got: %s", "nil", err.Error())
		t.Error(errorMessage)
	}
}

func TestFoodService_DeleteFood(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockFood := mocks.NewMockIFoodRepository(mockCtrl)
	testFood := &FoodService{foodRepo: mockFood}

	mockFood.EXPECT().DeleteFood(1).Return(nil)
	err := testFood.DeleteFood(1)
	if err != nil {
		errorMessage := fmt.Sprintf("Expected: %s, got: %s", "nil", err.Error())
		t.Error(errorMessage)
	}
}
