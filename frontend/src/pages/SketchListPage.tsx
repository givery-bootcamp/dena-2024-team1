import { useEffect } from "react";
import { useAtom } from "jotai";
import { AlertCircle } from "react-feather";

import { useAppDispatch, useAppSelector } from "~/shared/hooks";
import { APIService } from "~/shared/services";
import { SketchList } from "~/features/sketches/SketchList";
import { InfiniteCanvas } from "~/shared/components/InfiniteCanvas";
import { selectedSketchAtom } from "~/shared/store/Sketch";
import { SketchModal } from "~/features/sketches/SketchModal";

export function SketchListPage() {
  const { sketches } = useAppSelector((state) => state.sketches);
  const dispatch = useAppDispatch();
  const [selectedSketch, setSelectedSketch] = useAtom(selectedSketchAtom);

  useEffect(() => {
    dispatch(APIService.getSketches());
  }, [dispatch]);

  if (!sketches) return <p>Loading...</p>;
  if (sketches.length === 0) {
    return (
      <div className="flex h-screen-without-header flex-col items-center justify-center space-y-4 text-gray-200">
        <AlertCircle size={80} />
        <p className="text-4xl font-semibold">まだ投稿がありません</p>
      </div>
    );
  }
  return (
    <div className="h-screen-without-header w-screen">
      <InfiniteCanvas>
        <div className="p-10">
          <SketchList sketches={sketches} />
        </div>
      </InfiniteCanvas>
      {selectedSketch && (
        <SketchModal selectedSketch={selectedSketch} setSelectedSketch={setSelectedSketch} />
      )}
    </div>
  );
}
