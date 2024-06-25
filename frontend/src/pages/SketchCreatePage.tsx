import axios from "axios";
import { useRef, useState } from "react";

import { SketchHandWriter } from "~/features/sketches/SketchHandWriter";
import { Button } from "~/shared/components/Button";

export const SketchCreatePage = () => {
  const [imageUrl, setImageUrl] = useState<string>("");
  const canvasRef = useRef<HTMLCanvasElement>(null);

  const handleCanvasUpdate = (event: HTMLCanvasElement) => {
    setImageUrl(event.toDataURL("image/png"));
  };

  const postImageToApi = async () => {
    const canvas = canvasRef.current;
    if (canvas === null) return;
    const image = canvas.toDataURL("image/png");

    await axios.post("http://localhost:9000/images", { image });
  };

  return (
    <div>
      <SketchHandWriter canvas={canvasRef} onCanvasUpdate={handleCanvasUpdate} />
      <Button type="submit" onClick={postImageToApi}>APIに送る</Button>

      <div>
        <p>Preview</p>
        <img src={imageUrl} />
      </div>
    </div>
  );
};
