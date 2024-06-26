import { useEffect } from "react";

import { useAppDispatch, useAppSelector } from "~/shared/hooks";
import { APIService } from "~/shared/services";
import { SketchList } from "~/features/sketches/SketchList";
import { InfiniteCanvas } from "~/shared/components/InfiniteCanvas";

export function SketchListPage() {
  const { sketches } = useAppSelector((state) => state.sketches);
  const dispatch = useAppDispatch();

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
    </div>
  );
}
