export const downloadFile = (fileName: string, blob: Blob) => {
    const downloadUrl = window.URL.createObjectURL(blob)
    const link = document.createElement('a')
    link.href = downloadUrl
    link.setAttribute('download', fileName)
    document.body.appendChild(link)
    link.click()
    link.remove()
    window.URL.revokeObjectURL(downloadUrl)
}