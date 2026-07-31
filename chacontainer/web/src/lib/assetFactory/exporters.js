import { toPng, toCanvas } from 'html-to-image'
import { jsPDF } from 'jspdf'

function triggerDownload(href, fileName) {
  const link = document.createElement('a')
  link.href = href
  link.download = fileName
  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
}

// node: DOM element rendered at its real (unscaled) pixel size.
export async function exportNodeAsPng(node, fileName, pixelRatio = 2) {
  const dataUrl = await toPng(node, { pixelRatio, cacheBust: true })
  triggerDownload(dataUrl, fileName)
}

// nodes: array of DOM elements, one per PDF page, each already sized to pageSize (px).
// pageSize: 'A4' or [widthPx, heightPx] (used as the PDF page dimensions, matched 1:1 to node px).
// Las páginas se capturan como JPEG comprimido (no PNG): son diseños de
// fondo sólido + texto + gradientes, no fotografía, así que PNG sin comprimir
// infla el archivo a decenas de MB sin ganancia visual perceptible.
export async function exportNodesAsPdf(nodes, fileName, pageSize = 'A4') {
  let pdf
  for (let i = 0; i < nodes.length; i++) {
    const canvas = await toCanvas(nodes[i], { pixelRatio: 2, cacheBust: true })
    const imgData = canvas.toDataURL('image/jpeg', 0.9)

    if (!pdf) {
      const orientation = canvas.width >= canvas.height ? 'landscape' : 'portrait'
      const format = Array.isArray(pageSize) ? [pageSize[0], pageSize[1]] : pageSize
      pdf = new jsPDF({ orientation, unit: 'px', format, compress: true })
    } else {
      pdf.addPage()
    }

    const pageWidth = pdf.internal.pageSize.getWidth()
    const pageHeight = pdf.internal.pageSize.getHeight()
    pdf.addImage(imgData, 'JPEG', 0, 0, pageWidth, pageHeight)
  }
  pdf.save(fileName)
}

export function exportHtmlString(htmlString, fileName) {
  const blob = new Blob([htmlString], { type: 'text/html' })
  const url = URL.createObjectURL(blob)
  triggerDownload(url, fileName)
  URL.revokeObjectURL(url)
}
