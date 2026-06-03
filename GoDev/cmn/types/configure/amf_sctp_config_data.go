/** Copyright(C),2020-2022
* Author: zmj
* Date: 12/7/20 2:45 PM
* Description:
 */
package configure

type SctpParameter struct {
	Initmsg  SctpInitmsgConfig `yaml:"init msg"`
	Heatbeat SctpHeatbeatConfig
}

type SctpInitmsgConfig struct {
	NumOstreams    int `yaml:"num ostreams"`
	MaxInstreams   int `yaml:"max instreams"`
	MaxAttempts    int `yaml:"max attempts"`
	MaxInitTimeout int `yaml:"max init timeout"`
}

type SctpHeatbeatConfig struct {
	Interval   int `yaml:"interval"`
	PathMaxRXT int `yaml:"path max retry"`
}
