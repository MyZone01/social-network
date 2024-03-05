import fs from "fs";
import path from "path";

const folderPath = "./";
const relativeFolderPath = path.relative(process.cwd(), folderPath);

export default () => {
  const files = fs
    .readdirSync(folderPath)
    .filter((file) => file.match(/\.(js)$/i));

  const filesPaths = files.map(
    (fileName) => `/${relativeFolderPath}/${fileName}`
  );

  return {
    filesPaths,
  };
};
