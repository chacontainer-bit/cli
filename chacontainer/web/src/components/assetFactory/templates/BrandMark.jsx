import { Package } from 'lucide-react'
import { BRAND } from '../../../lib/assetFactory/brand'

export default function BrandMark({ size = 20, light = false, textSize = 15 }) {
  const color = light ? '#ffffff' : BRAND.color
  return (
    <div style={{ display: 'flex', alignItems: 'center', gap: 8 }}>
      <Package size={size} color={color} strokeWidth={2.2} />
      <span
        style={{
          color,
          fontWeight: 800,
          letterSpacing: '0.14em',
          fontSize: textSize,
          fontFamily: BRAND.fontFamily,
        }}
      >
        {BRAND.name}
      </span>
    </div>
  )
}
