import { useRef, type FC } from "react";
import {
  ReactInfiniteCanvas,
  type ReactInfiniteCanvasHandle,
  type ReactInfiniteCanvasProps,
} from "react-infinite-canvas";

const config = {
  scale: 0.5,
  backgroundColor: "white",
};

export const InfiniteCanvas: FC<ReactInfiniteCanvasProps> = (props) => {
  const canvasRef = useRef<ReactInfiniteCanvasHandle>();

  return (
    <ReactInfiniteCanvas
      ref={canvasRef}
      onCanvasMount={(canvasFunc) => {
        canvasFunc.fitContentToView({ scale: config.scale });
      }}
      backgroundConfig={{
        elementColor: config.backgroundColor,
      }}
      {...props}
    />
  );
};
