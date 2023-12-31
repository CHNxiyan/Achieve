export const SubtitleLocation = {
  TopLeft: 7,
  TopCenter: 8,
  TopRight: 9,
  BottomLeft: 1,
  BottomCenter: 2,
  BottomRight: 3,
}
const getLocationName = (location: number) =>
  Object.entries(SubtitleLocation)
    .filter(([, value]) => value === location)
    .map(([name]) => name)
    .shift()
export enum SubtitleSize {
  VerySmall = 1,
  Small,
  Medium,
  Large,
  VeryLarge,
}
export interface SubtitleConverterConfig {
  title: string
  width: number
  height: number
  color: string
  size: number
  opacity: number
  location: number
  boxPadding: number
  boxMargin: number
}
export interface SubtitleItem {
  from: number
  to: number
  location: number
  content: string
}
export class SubtitleConverter {
  static defaultConfig: SubtitleConverterConfig
  config: SubtitleConverterConfig
  constructor(config: {
    title: string
    width: number
    height: number
    color?: string
    size?: number
    opacity?: number
    location?: number
    boxPadding?: number
    boxMargin?: number
  }) {
    this.config = Object.assign(SubtitleConverter.defaultConfig, config)
  }
  private async getAssMeta() {
    const { convertHexColorForStyle } = await import('@/components/video/ass-utils')
    const color = convertHexColorForStyle(this.config.color) // 前景色貌似不受字幕设置里的透明度控制
    const background = convertHexColorForStyle('#000000', this.config.opacity)
    const fontStyles = []
    const fontSize = ((10 * (this.config.size - 3) + 48) * this.config.height) / 720
    console.log(fontSize)
    for (const [name, value] of Object.entries(SubtitleLocation)) {
      fontStyles.push(
        `Style: ${name},微软雅黑,${fontSize},${color},${color},${background},${background},0,0,0,0,100,100,0,0,3,${this.config.boxPadding},0,${value},${this.config.boxMargin},${this.config.boxMargin},${this.config.boxMargin},0`,
      )
    }
    return `
[Script Info]
; Script generated by Bilibili Evolved Danmaku Converter
; https://github.com/the1812/Bilibili-Evolved/
Title: ${this.config.title}
ScriptType: v4.00+
PlayResX: ${this.config.width}
PlayResY: ${this.config.height}
Timer: 10.0000
WrapStyle: 0
ScaledBorderAndShadow: no

[V4+ Styles]
Format: Name, Fontname, Fontsize, PrimaryColour, SecondaryColour, OutlineColour, BackColour, Bold, Italic, Underline, StrikeOut, ScaleX, ScaleY, Spacing, Angle, BorderStyle, Outline, Shadow, Alignment, MarginL, MarginR, MarginV, Encoding
${fontStyles.join('\n')}

[Events]
Format: Layer, Start, End, Style, Name, MarginL, MarginR, MarginV, Effect, Text`.trim()
  }
  async convertToAss(subtitles: SubtitleItem[]) {
    const { convertTimeByEndTime, normalizeContent } = await import('@/components/video/ass-utils')
    const meta = await this.getAssMeta()
    return `${meta}\n${subtitles
      .map(s => {
        const [startTime, endTime] = convertTimeByEndTime(s.from, s.to)
        const dialogue = `Dialogue: 0,${startTime},${endTime},${getLocationName(
          this.config.location,
        )},,0,0,0,,${normalizeContent(s.content)}`
        return dialogue
      })
      .join('\n')}`
  }
}
SubtitleConverter.defaultConfig = {
  title: '',
  color: '#ffffff',
  width: 1920,
  height: 1080,
  size: SubtitleSize.Medium,
  opacity: 0.5,
  location: SubtitleLocation.BottomCenter,
  boxPadding: 1,
  boxMargin: 32,
}
