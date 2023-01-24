package usecase

type MathUseCase struct {
	service MathService
}

func NewMathUC(ma MathService) *MathUseCase {
	return &MathUseCase{service: ma}
}

func (uc *MathUseCase) Calculate(text string) (float64, error) {
	return uc.service.Calculate(text)
}
