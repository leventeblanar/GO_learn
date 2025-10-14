package units

var LengthUnits = map[string]float64{
    "millimeter": 0.001,
    "centimeter": 0.01, 
    "meter": 1.0,       
    "kilometer": 1000.0,
    "inch": 0.0254,     
    "foot": 0.3048,     
    "yard": 0.9144,     
    "mile": 1609.34,    
}

var WeightUnits = map[string]float64{
    "milligram": 1e-6,   
    "gram": 1e-3,        
    "kilogram": 1.0,     
    "ounce": 0.028349523125,
    "pound": 0.45359237,    
}

var Aliases = map[string]string{
	//length
    "mm": "millimeter", "millimeter": "millimeter",
    "cm": "centimeter", "centimeter": "centimeter",
    "m": "meter", "meter": "meter",
    "km": "kilometer", "kilometer": "kilometer",
    "in": "inch", "inch": "inch",
    "ft": "foot", "foot": "foot",
    "yd": "yard", "yard": "yard",
    "mi": "mile", "mile": "mile",
    // weight
    "mg": "milligram", "milligram": "milligram",
    "g": "gram", "gram": "gram",
    "kg": "kilogram", "kilogram": "kilogram",
    "oz": "ounce", "ounce": "ounce",
    "lb": "pound", "lbs": "pound", "pound": "pound",
    // temperature
    "C": "celsius", "°c": "celsius", "celsius": "celsius",
    "F": "fahrenheit", "°f": "fahrenheit", "fahrenheit": "fahrenheit",
    "K": "kelvin", "kelvin": "kelvin",
}