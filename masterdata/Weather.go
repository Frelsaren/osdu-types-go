package masterdata

import "time"

// Meteorological readings for the defined time period on an operations report
type Weather struct {
	// Name of company that supplied the weather data.                                                    
	Agency                                                                                      *string   `json:"Agency,omitempty"`
	// Atmospheric pressure.                                                                              
	BarometricPressure                                                                          *float64  `json:"BarometricPressure,omitempty"`
	// The Beaufort wind force scale is a system used to estimate and report wind speeds when no          
	// measuring apparatus is available. It was invented in the early 19th century by Admiral             
	// Sir Francis Beaufort of the British Navy as a way to interpret winds from conditions.              
	// Values range from 0 (calm) to 12 (hurricane force).                                                
	BeaufortScaleNumber                                                                         *float64  `json:"BeaufortScaleNumber,omitempty"`
	// Height of cloud cover.                                                                             
	CloudCeiling                                                                                *float64  `json:"CloudCeiling,omitempty"`
	// Description of cloud cover.                                                                        
	CloudCover                                                                                  *string   `json:"CloudCover,omitempty"`
	// Comments and remarks                                                                               
	Comments                                                                                    *string   `json:"Comments,omitempty"`
	// The speed of the ocean current.                                                                    
	CurrentSeaSpeed                                                                             *float64  `json:"CurrentSeaSpeed,omitempty"`
	// Date and time that weather was observed                                                            
	DateTime                                                                                    time.Time `json:"DateTime"`
	// Amount of precipitation.                                                                           
	PrecipitationAmount                                                                         *float64  `json:"PrecipitationAmount,omitempty"`
	// Azimuth of current.                                                                                
	SeaCurrentDirection                                                                         *float64  `json:"SeaCurrentDirection,omitempty"`
	// Sea temperature.                                                                                   
	SeaTemp                                                                                     *float64  `json:"SeaTemp,omitempty"`
	// Maximum temperature above ground.                                                                  
	SurfaceTempMax                                                                              *float64  `json:"SurfaceTempMax,omitempty"`
	// Minimum temperature above ground. Temperature of the atmosphere.                                   
	SurfaceTempMin                                                                              *float64  `json:"SurfaceTempMin,omitempty"`
	// Horizontal visibility.                                                                             
	Visibility                                                                                  *float64  `json:"Visibility,omitempty"`
	// The direction from which the waves are coming, measured from true north.                           
	WaveDirection                                                                               *float64  `json:"WaveDirection,omitempty"`
	// Average height of the sea waves.                                                                   
	WaveHeightAverage                                                                           *float64  `json:"WaveHeightAverage,omitempty"`
	// The maximum wave height.                                                                           
	WaveHeightMax                                                                               *float64  `json:"WaveHeightMax,omitempty"`
	// DEPRECATED: The elapsed time between the passing of two wave tops.                                 
	WavePeriod                                                                                  *string   `json:"WavePeriod,omitempty"`
	// The elapsed time between the passing of two wave tops.                                             
	WavePeriodDuration                                                                          *float64  `json:"WavePeriodDuration,omitempty"`
	// An average of the higher 1/3 of the wave heights passing during a sample period                    
	// (typically 20 to 30 minutes).                                                                      
	WaveSignificant                                                                             *float64  `json:"WaveSignificant,omitempty"`
	// Type of weather.                                                                                   
	WeatherType                                                                                 *string   `json:"WeatherType,omitempty"`
	// A measure of the combined chilling effect of wind and low temperature on living things,            
	// also named chill factor, e.g., according to the US weather service table, an air                   
	// temperature of 30 degF with a 10 mph corresponds to a windchill of 22 degF.                        
	WindChillTemp                                                                               *float64  `json:"WindChillTemp,omitempty"`
	// The direction from which the wind is blowing, measured from true north.                            
	WindDirection                                                                               *float64  `json:"WindDirection,omitempty"`
	// Wind speed.                                                                                        
	WindSpeed                                                                                   *float64  `json:"WindSpeed,omitempty"`
}
