import { Csound } from "@csound/browser";
import song from "../csound/song.csd";

export const playSong = async () => {
  const csoundObj = await Csound({
    useWorker: false,
    useSPN: true,
    useSAB: false
  });

  if (!csoundObj) return;
  await csoundObj.compileCsdText(song);

  return csoundObj.start();
};
