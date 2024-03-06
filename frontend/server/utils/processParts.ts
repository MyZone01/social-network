// Function to handle file part
const handleFilePart = async (part) => {
  if (!part.type.startsWith('image')) {
    throw new Error('File is not an image')
  }
  return part
}

// Function to handle data part
const handleDataPart = (part) => {
  const data = Buffer.from(part.data).toString()
  return JSON.parse(data)
}

// Function to process parts
export const processParts = async (reader) => {
  let file, jsonData
  for await (const part of reader) {
    if (part.name === 'file') {
      if (part.filename) {
        file = await handleFilePart(part)
      }
      continue
    } else if (part.name === 'data') {
      jsonData = handleDataPart(part)
    }
  }
  return { file, jsonData }
}