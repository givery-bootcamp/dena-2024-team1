import { useEffect,Fragment } from "react";

import { useAppDispatch, useAppSelector } from "~/shared/hooks";
import { APIService } from "~/shared/services";
import { Container } from "~/shared/components/Container";
import { SketchItem } from "~/features/sketches/SketchItem";
import { BorderLine } from "~/shared/components/BorderLine";
import { Sketch } from "~/generated";

export function ImageListPage() {
  const { sketches } = useAppSelector((state) => state.sketches);
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(APIService.getSketches());
  }, [dispatch]);
  console.log(sketches);
  if (!sketches) {
    return <Container>loading...</Container>;
  }
  // 3つの要素ごとに配列を分割する関数
  const chunkArray = (array: Sketch[], size: number): Sketch[][] => {
    return array.reduce((acc: Sketch[][], _, index: number) => 
      index % size === 0 ? [...acc, array.slice(index, index + size)] : acc, []);
  };


  // sketchesを3つの要素ごとに分割
  const sketchesChunks = chunkArray(sketches, 3);
  
  return (
    <Container className='mt-10'>
      <h1 className='text-2xl font-bold'>画像一覧</h1>
      {sketchesChunks.map((chunk, chunkIndex) => (
        <Fragment key={chunkIndex}>
          <div className="flex justify-between"> {/* ここでflexboxを使って横並びにします */}
            {chunk.map((sketch) => (
              <Fragment key={sketch.id}>
                <SketchItem sketch={sketch}/>
              </Fragment>
            ))}
          </div>
          {chunkIndex < sketchesChunks.length - 1 && <BorderLine/>}
        </Fragment>
      ))}
    </Container>
  );
}
