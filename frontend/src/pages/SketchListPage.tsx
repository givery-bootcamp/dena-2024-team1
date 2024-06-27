import { useEffect } from "react";
import { useAtom } from "jotai";

import { Modal } from "~/shared/components/Modal";
import { useAppDispatch, useAppSelector } from "~/shared/hooks";
import { APIService } from "~/shared/services";
import { SketchList } from "~/features/sketches/SketchList";
import { InfiniteCanvas } from "~/shared/components/InfiniteCanvas";
import { selectedSketchUrlAtom } from "~/shared/store/Sketch";

export function SketchListPage() {
  const { sketches } = useAppSelector((state) => state.sketches);
  const dispatch = useAppDispatch();
  const [selectedSketchUrl, setSelectedSketchUrl] = useAtom(selectedSketchUrlAtom);
  
  useEffect(() => {
    dispatch(APIService.getSketches());
  }, [dispatch]);

  if (!sketches) return <p>Loading...</p>;
  return (
    <div className="h-screen w-screen">
      <InfiniteCanvas>
        <div className="p-10">
          <SketchList sketches={sketches} />
        </div>
      </InfiniteCanvas>
      {selectedSketchUrl && (
        <Modal isOpen={true} onClose={() => setSelectedSketchUrl(null)}>
          <img src={selectedSketchUrl} alt={selectedSketchUrl} />
        </Modal>
      )}
    </div>
  );
}
