import { formatDateTime } from "~/shared/utils";
import { Sketch } from "~/generated";

type SketchItemProps = {sketch: Sketch};

export function SketchItem({ sketch }: SketchItemProps) {
  return(
    <div className="flex flex-col items-center border border-solid border-gray-200 py-4">
      <img src={sketch.imageUrl} alt={sketch.imageUrl} className="w-1/2"/>
      <div>
        <div className='text-sm font-semibold text-gray-100'>
          {sketch.userName}
        </div>
        <div className='text-sm text-gray-200'>
          更新日時: {formatDateTime(sketch.updatedAt)}
        </div>
      </div>
    </div>
  );
}
