import { useEffect } from "react";

import { useAppDispatch, useAppSelector } from "~/shared/hooks";
import { APIService } from "~/shared/services";
import { Container } from "~/shared/components/Container";
import { SketchList } from "~/features/sketches/SketchList";

export function SketchListPage() {
  const { sketches } = useAppSelector((state) => state.sketches);
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(APIService.getSketches());
  }, [dispatch]);

  if (!sketches) return <p>Loading...</p>;
  return (
    <Container className='mt-10'>
      <SketchList sketches={sketches} />
    </Container>
  );
}
