import { useEffect, useRef } from "react";
import { ReactInfiniteCanvas, type ReactInfiniteCanvasHandle } from "react-infinite-canvas";

import { useAppDispatch, useAppSelector } from "~/shared/hooks";
import { APIService } from "~/shared/services";
import { SketchList } from "~/features/sketches/SketchList";

export function SketchListPage() {
  const { sketches } = useAppSelector((state) => state.sketches);
  const dispatch = useAppDispatch();
  const canvasRef = useRef<ReactInfiniteCanvasHandle>();

  useEffect(() => {
    dispatch(APIService.getSketches());
  }, [dispatch]);

  if (!sketches) return <p>Loading...</p>;
  return (
    <div className="h-screen w-screen">
      <ReactInfiniteCanvas
        ref={canvasRef}
        onCanvasMount={(canvasFunc) => {
          canvasFunc.fitContentToView({ scale: 0.5 });
        }}
        backgroundConfig={{
          elementColor: "white",
        }}
      >
        <SketchList sketches={sketches} />
      </ReactInfiniteCanvas>
    </div>
  );
}
