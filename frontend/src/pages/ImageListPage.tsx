import { useEffect,Fragment } from "react";

import { useAppDispatch, useAppSelector } from "~/shared/hooks";
import { APIService } from "~/shared/services";
import { Container } from "~/shared/components/Container";
import { SketchItem } from "~/features/sketches/SketchItem";
import { BorderLine } from "~/shared/components/BorderLine";

export function ImageListPage() {
  const { sketches } = useAppSelector((state) => state.sketches);
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(APIService.getSketches());
  }, [dispatch]);
  console.log(sketches);
  return (
    <Container className='mt-10'>
      <h1 className='text-2xl font-bold'>画像一覧</h1>
      {sketches?.map((sketch,index)=>(
        <Fragment key={index}>
          <SketchItem sketch={sketch}/>
          {index === sketches.length - 1 ? null : <BorderLine/>}
        </Fragment>
      ))}
    </Container>
  );
}