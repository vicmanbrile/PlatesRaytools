package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// PlateSkeleton es una estructura separada para mejorar la legibilidad
type PlateSkeleton struct {
	Text           string `xml:",chardata"`
	CommonProperty struct {
		Text          string `xml:",chardata"`
		Selected      string `xml:"selected,attr"`
		Layer         string `xml:"layer,attr"`
		Tlayer        string `xml:"tlayer,attr"`
		Tlayerb       string `xml:"tlayerb,attr"`
		Note          string `xml:"note,attr"`
		PartClr       string `xml:"PartClr,attr"`
		PartName      string `xml:"partName,attr"`
		PartNsetInfo  string `xml:"partNsetInfo,attr"`
		PartGuid      string `xml:"partGuid,attr"`
		PartEdgeGuids string `xml:"partEdgeGuids,attr"`
	} `xml:"CommonProperty"`
	Shapes string `xml:"shapes"`
}

// PlateNodeStruct es una estructura separada para PlateNode
type PlateNodeStruct struct {
	Text              string `xml:",chardata"`
	PlateNodeInfo struct {
		Text            string `xml:",chardata"`
		Guid            string `xml:"Guid,attr"`
		Guid00          string `xml:"Guid00,attr"`
		PlateName       string `xml:"PlateName,attr"`
		IsRectangle     string `xml:"IsRectangle,attr"`
		PlateProcessNum string `xml:"PlateProcessNum,attr"`
		TotalNum        string `xml:"TotalNum,attr"`
		Width           string `xml:"Width,attr"`
		Height          string `xml:"Height,attr"`
		UtilizationRate string `xml:"UtilizationRate,attr"`
		PlateGrade      string `xml:"PlateGrade,attr"`
	} `xml:"PlateNodeInfo"`
	Shapes            string `xml:"shapes"`
	PlateProfilesInfo string `xml:"PlateProfilesInfo"`
	PlateBoardPnts struct {
		Text          string `xml:",chardata"`
		PlateBoardPnt1s struct {
			Text          string `xml:",chardata"`
			PlateBoardPnt []struct {
				Text string `xml:",chardata"`
				X    string `xml:"x,attr"`
				Y    string `xml:"y,attr"`
				Bul  string `xml:"bul,attr"`
			} `xml:"PlateBoardPnt"`
		} `xml:"PlateBoardPnt1s"`
	} `xml:"PlateBoardPnts"`
	// PlateSkeleton ahora es un puntero para manejar su ausencia
	PlateSkeleton *PlateSkeleton `xml:"PlateSkeleton,omitempty"` 
}

// StopPosStruct es una estructura separada para StopPos
type StopPosStruct struct {
	Text                    string `xml:",chardata"`
	StopPosType             string `xml:"StopPosType,attr"`
	X                       string `xml:"x,attr"`
	Y                       string `xml:"y,attr"`
	StopPosExcludeUnprocess string `xml:"StopPosExcludeUnprocess,attr"`
	StopPosExcludeShapeTool string `xml:"StopPosExcludeShapeTool,attr"`
	StopPosExcludeBoard     string `xml:"StopPosExcludeBoard,attr"`
}

// NestParamStruct es una estructura separada para NestParam
type NestParamStruct struct {
	Text                    string `xml:",chardata"`
	NestAngle               string `xml:"nestAngle,attr"`
	AutoCommonEdge          string `xml:"autoCommonEdge,attr"`
	RealNest                string `xml:"realNest,attr"`
	DisCountInLeads         string `xml:"disCountInLeads,attr"`
	RealShape               string `xml:"RealShape,attr"`
	MirrorType              string `xml:"mirrorType,attr"`
	StartCornerType         string `xml:"startCornerType,attr"`
	NestDir                 string `xml:"nestDir,attr"`
	BestNestDir             string `xml:"BestNestDir,attr"`
	HoleNes                 string `xml:"holeNes,attr"`
	PreferHoleNest          string `xml:"preferHoleNest,attr"`
	Resolution              string `xml:"resolution,attr"`
	InRotateAngle           string `xml:"inRotateAngle,attr"`
	AutoComb                string `xml:"AutoComb,attr"`
	NestCombEdge            string `xml:"NestCombEdge,attr"`
	NestCombEdgeProximity   string `xml:"NestCombEdgeProximity,attr"`
	NestCombType            string `xml:"NestCombType,attr"`
	RealNestEngine2         string `xml:"realNestEngine2,attr"`
	ShowlNestEngine2        string `xml:"showlNestEngine2,attr"`
	NestAllCloseShapes      string `xml:"nestAllCloseShapes,attr"`
	NestAllGrid             string `xml:"nestAllGrid,attr"`
	NestNextGen             string `xml:"nestNextGen,attr"`
	BClearAlreadyNest       string `xml:"bClearAlreadyNest,attr"`
	FlatEndParam            string `xml:"FlatEndParam,attr"`
	NestEngineType          string `xml:"NestEngineType,attr"`
	EnableNestYT            string `xml:"enableNestYT,attr"`
	NestCombEdgeYT          string `xml:"NestCombEdgeYT,attr"`
	NestAllCombEdgeYT       string `xml:"NestAllCombEdgeYT,attr"`
	NestMaxCombEdgeNumYT    string `xml:"NestMaxCombEdgeNumYT,attr"`
	NestCombEdgeSamePartYT  string `xml:"NestCombEdgeSamePartYT,attr"`
	NestEnableTimeYT        string `xml:"NestEnableTimeYT,attr"`
	NestTimeYT              string `xml:"NestTimeYT,attr"`
	NestDis                 string `xml:"nestDis,attr"`
	NestDisBoard            string `xml:"nestDisBoard,attr"`
	EnableMinComnEdgeLen    string `xml:"EnableMinComnEdgeLen,attr"`
	MinComnEdgeLen          string `xml:"MinComnEdgeLen,attr"`
}


// Definición final de Root con punteros para campos opcionales
type Root struct {
	XMLName xml.Name `xml:"root"`
	Text    string   `xml:",chardata"`
	V1      string   `xml:"V1,attr,omitempty"` // Campo que faltaba/opcional
	V2      string   `xml:"V2,attr"`
	Unit    string   `xml:"unit,attr"`
	ShowModel string `xml:"ShowModel,attr"`
	ShowItemGuid string `xml:"ShowItemGuid,attr"` // Variable: no se pone por defecto
	ViewX   string   `xml:"ViewX,attr"`           // Variable: no se pone por defecto
	ViewY   string   `xml:"ViewY,attr"`           // Variable: no se pone por defecto
	ViewScale string `xml:"ViewScale,attr"`       // Variable: no se pone por defecto
	
	// StopPos y NestParam ahora son punteros para ser opcionales (omitempty)
	StopPos   *StopPosStruct   `xml:"StopPos,omitempty"` 
	NestParam *NestParamStruct `xml:"NestParam,omitempty"` 

	PlateNode []PlateNodeStruct `xml:"PlateNode"`
}

// createDefaultRoot inicializa la estructura Root con valores por defecto.
func createDefaultRoot() Root {
	return Root{
		// Atributos con valores por defecto consistentes o que deben existir
		// V1 se deja vacío (si falta) o se le asigna un valor común
		V1: "1.0.18.24443", // Valor del archivo 4510x1500.xml
		V2: "1.0.0.31202", // Valor más común
		Unit: "4", 
		ShowModel: "3", // Valor más común

		// Atributos inherentemente variables (GUID, Coordenadas) deben ser inicializados
		// a "0" o un valor de inicio, pero no son "por defecto" en el sentido de "constante"
		ShowItemGuid: "00000000-0000-0000-0000-000000000000",
		ViewX: "0.0",
		ViewY: "0.0",
		ViewScale: "1.0",
		
		// StopPos y NestParam se dejan como nil por defecto.
		// El código principal puede inicializarlos si son necesarios.
		StopPos: nil,
		NestParam: nil,

		PlateNode: []PlateNodeStruct{
			{
				PlateNodeInfo: PlateNodeStruct.PlateNodeInfo{
					Guid: "00000000-0000-0000-0000-000000000000",
					Guid00: "00000000-0000-0000-0000-000000000000",
					PlateName: "Default Plate",
					IsRectangle: "false",
					PlateProcessNum: "0",
					TotalNum: "1",
					Width: "2000",
					Height: "1000",
					UtilizationRate: "0",
					PlateGrade: "2",
				},
				Shapes:            "",
				PlateProfilesInfo: "",
				PlateBoardPnts: PlateNodeStruct.PlateBoardPnts{
					PlateBoardPnt1s: PlateNodeStruct.PlateBoardPnts.PlateBoardPnt1s{
						PlateBoardPnt: []struct {
							Text string `xml:",chardata"`
							X    string `xml:"x,attr"`
							Y    string `xml:"y,attr"`
							Bul  string `xml:"bul,attr"`
						}{
							{X: "0", Y: "0", Bul: "0"},
							{X: "2000", Y: "0", Bul: "0"},
							{X: "2000", Y: "1000", Bul: "0"},
							{X: "0", Y: "1000", Bul: "0"},
							{X: "0", Y: "0", Bul: "0"},
						},
					},
				},
				PlateSkeleton: nil, // Se deja nil por defecto
			},
		},
	}
}

// createDefaultStopPos crea una estructura StopPos por defecto
func createDefaultStopPos() *StopPosStruct {
	return &StopPosStruct{
		StopPosType: "6",
		X: "0.0",
		Y: "0.0",
		StopPosExcludeUnprocess: "false",
		StopPosExcludeShapeTool: "false",
		StopPosExcludeBoard: "true",
	}
}

// createDefaultNestParam crea una estructura NestParam por defecto
func createDefaultNestParam() *NestParamStruct {
	return &NestParamStruct{
		NestAngle: "-0.785398163397448",
		AutoCommonEdge: "false",
		RealNest: "-1",
		DisCountInLeads: "false",
		RealShape: "-1",
		MirrorType: "0",
		StartCornerType: "0",
		NestDir: "1",
		BestNestDir: "false",
		HoleNes: "false",
		PreferHoleNest: "false",
		Resolution: "0.5",
		InRotateAngle: "90",
		AutoComb: "false",
		NestCombEdge: "false",
		NestCombEdgeProximity: "false",
		NestCombType: "1",
		RealNestEngine2: "true",
		ShowlNestEngine2: "true",
		NestAllCloseShapes: "false",
		NestAllGrid: "true",
		NestNextGen: "true",
		BClearAlreadyNest: "true",
		FlatEndParam: "true",
		NestEngineType: "0",
		EnableNestYT: "false",
		NestCombEdgeYT: "false",
		NestAllCombEdgeYT: "false",
		NestMaxCombEdgeNumYT: "2",
		NestCombEdgeSamePartYT: "false",
		NestEnableTimeYT: "false",
		NestTimeYT: "1",
		NestDis: "3",
		NestDisBoard: "1",
		EnableMinComnEdgeLen: "false",
		MinComnEdgeLen: "1",
	}
}


func main() {
	// 1. Crear la estructura con valores por defecto
	defaultRoot := createDefaultRoot()

	fmt.Println("--- XML con Valores por Defecto (StopPos y NestParam NULOS) ---")
	printXML(defaultRoot)

	// 2. Ejemplo para cambiar los valores por defecto (ejemplo de inicialización de StopPos y NestParam)
	// Inicializar y Modificar StopPos y NestParam
	defaultRoot.StopPos = createDefaultStopPos()
	defaultRoot.NestParam = createDefaultNestParam()
	
	defaultRoot.Unit = "inch"
	defaultRoot.ViewScale = "2.5"

	defaultRoot.StopPos.X = "150.0"
	defaultRoot.StopPos.Y = "200.0"

	defaultRoot.NestParam.RealNest = "0"
	defaultRoot.NestParam.NestAngle = "90"
	defaultRoot.NestParam.Resolution = "0.5"

	// Modificar valores en el primer PlateNode
	if len(defaultRoot.PlateNode) > 0 {
		defaultRoot.PlateNode[0].PlateNodeInfo.PlateName = "Custom Plate X"
		defaultRoot.PlateNode[0].PlateNodeInfo.Width = "3500"
	}

	fmt.Println("\n--- XML con Valores Modificados (StopPos y NestParam PRESENTES) ---")
	printXML(defaultRoot)
}

// printXML serializa e imprime la estructura Root como XML
func printXML(r Root) {
	// Añadir encabezado XML
	output, err := xml.MarshalIndent(r, "", "    ")
	if err != nil {
		fmt.Printf("Error al generar XML: %v\n", err)
		return
	}

	// El paquete xml.MarshalIndent no incluye la declaración XML por defecto.
	// La agregamos manualmente para un XML más estándar.
	finalXML := xml.Header + string(output)
	
	fmt.Println(finalXML)
	
	// Escribir a un archivo para verificación
	file, err := os.Create("output.xml")
	if err == nil {
	    file.Write([]byte(finalXML))
	    file.Close()
	}
}