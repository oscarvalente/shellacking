package shell_geo

func TranslateBy(x float64, y float64, state *[][][]float64) {
	t := CreateTranslationMatrix2D(x, y)
	*state = append(*state, t)
}

func ScaleBy(factorInX float64, factorInY float64, state *[][][]float64) {
	t := CreateScaleMatrix2D(factorInX, factorInY)
	*state = append(*state, t)
}

func RotateBy(angleDegrees float64, state *[][][]float64) {
	t := CreateRotationMatrix2D(angleDegrees)
	*state = append(*state, t)
}


