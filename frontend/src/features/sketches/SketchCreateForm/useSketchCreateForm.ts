import { RefObject } from "react";
import { useNavigate } from "react-router-dom";

import { sketchApi } from "~/shared/services/API";

type Args = {
  canvasRef: RefObject<HTMLCanvasElement>
}

export const useSketchCreateForm = ({ canvasRef }: Args) => {
  const navigate = useNavigate();

  const convertCanvasToFile = async (): Promise<File> => {
    const canvas = canvasRef.current;
    if (canvas === null) {
      throw new Error("canvasが存在しません");
    }

    const image = canvas.toDataURL("image/png");

    const blob = await fetch(image).then((res) => res.blob());
    const file = new File([blob], "sketch.png", { type: "image/png" });

    return file;
  };

  const createSketch = async () => {
    const showAlert = () => window.alert("投稿に失敗しました");
    const file = await convertCanvasToFile();
    try {
      const result = await sketchApi.postSketch(file);

      if (result.status === 201) {
        navigate("/");
      } else {
        showAlert();
      }

    } catch (error) {
      showAlert();
    }
  };

  return { createSketch };
};
