package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// --- DEFINICIÓN DE TIPOS SECUNDARIOS (PARA CLARIDAD Y FACILIDAD DE INICIALIZACIÓN) ---

// PlateNodeInfoStruct ahora es un tipo con nombre
type PlateNodeInfoStruct struct {
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
}

// PlateBoardPntStruct para los puntos individuales
type PlateBoardPntStruct struct {
	Text string `xml:",chardata"`
	X    string `xml:"x,attr"`
	Y    string `xml:"y,attr"`
	Bul  string `xml:"bul,attr"`
}

// PlateBoardPnt1sStruct para el contenedor de puntos
type PlateBoardPnt1sStruct struct {
	Text          string `xml:",chardata"`
	PlateBoardPnt []PlateBoardPntStruct `xml:"PlateBoardPnt"`
}

// PlateBoardPntsStruct para el contenedor principal de los puntos
type PlateBoardPntsStruct struct {
	Text          string `xml:",chardata"`
	PlateBoardPnt1s PlateBoardPnt1sStruct `xml:"PlateBoardPnt1s"`
}

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

// PlateNodeStruct ahora usa los tipos con nombre
type PlateNodeStruct struct {
	Text              string `xml:",chardata"`
	PlateNodeInfo     PlateNodeInfoStruct `xml:"PlateNodeInfo"`
	Shapes            string `xml:"shapes"`
	PlateProfilesInfo string `xml:"PlateProfilesInfo"`
	PlateBoardPnts    PlateBoardPntsStruct `xml:"PlateBoardPnts"`
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
	V1      string   `xml:"V1,attr,omitempty"` 
	V2      string   `xml:"V2,attr"`
	Unit    string   `xml:"unit,attr"`
	ShowModel string `xml:"ShowModel,attr"`
	ShowItemGuid string `xml:"ShowItemGuid,attr"` 
	ViewX   string   `xml:"ViewX,attr"`           
	ViewY   string   `xml:"ViewY,attr"`           
	ViewScale string `xml:"ViewScale,attr"`       
	
	StopPos   *StopPosStruct   `xml:"StopPos,omitempty"` 
	NestParam *NestParamStruct `xml:"NestParam,omitempty"` 

	PlateNode []PlateNodeStruct `xml:"PlateNode"`
}

// createDefaultRoot inicializa la estructura Root con valores por defecto.
func createDefaultRoot() Root {
	return Root{
		V1: "1.0.18.24443", 
		V2: "1.0.0.31202", 
		Unit: "4", 
		ShowModel: "3", 

		ShowItemGuid: "00000000-0000-0000-0000-000000000000",
		ViewX: "0.0",
		ViewY: "0.0",
		ViewScale: "1.0",
		
		StopPos: nil,
		NestParam: nil,

		PlateNode: []PlateNodeStruct{
			{
				// CORRECTO: Usamos el tipo con nombre PlateNodeInfoStruct
				PlateNodeInfo: PlateNodeInfoStruct{
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
				// CORRECTO: Usamos el tipo con nombre PlateBoardPntsStruct
				PlateBoardPnts: PlateBoardPntsStruct{
					// CORRECTO: Usamos el tipo con nombre PlateBoardPnt1sStruct
					PlateBoardPnt1s: PlateBoardPnt1sStruct{
						// CORRECTO: Usamos un slice del tipo con nombre PlateBoardPntStruct
						PlateBoardPnt: []PlateBoardPntStruct{
							{X: "0", Y: "0", Bul: "0"},
							{X: "2000", Y: "0", Bul: "0"},
							{X: "2000", Y: "1000", Bul: "0"},
							{X: "0", Y: "1000", Bul: "0"},
							{X: "0", Y: "0", Bul: "0"},
						},
					},
				},
				PlateSkeleton: nil, 
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
	output, err := xml.MarshalIndent(r, "", "    ")
	if err != nil {
		fmt.Printf("Error al generar XML: %v\n", err)
		return
	}

	finalXML := xml.Header + string(output)
	
	fmt.Println(finalXML)
	
	file, err := os.Create("output.xml")
	if err == nil {
	    file.Write([]byte(finalXML))
	    file.Close()
	}
}
