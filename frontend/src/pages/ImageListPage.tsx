import { useEffect,Fragment } from "react";

import { useAppDispatch, useAppSelector } from "~/shared/hooks";
import { APIService } from "~/shared/services";
import { Container } from "~/shared/components/Container";
import { SketchItem } from "~/features/sketches/SketchItem";

export function ImageListPage() {
  const { sketches } = useAppSelector((state) => state.sketches);
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(APIService.getSketches());
  }, [dispatch]);

  
  return (
    <Container className='mt-10'>
      <div className='grid grid-cols-3 gap-10'>
        {sketches?.map((sketch,index) => (
          <Fragment key={index}>
            <SketchItem sketch={sketch}/>
          </Fragment>
      ))}
      </div>
    </Container>
  );
}
