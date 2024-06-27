import { useState }from "react";

import { Modal } from "~/shared/components/Modal";
import { formatDateTime } from "~/shared/utils";
import { Sketch } from "~/generated";

type SketchItemProps = {sketch: Sketch};

export function SketchItem({ sketch }: SketchItemProps) {
  // モーダルの表示状態を管理する状態変数
  const [isModalOpen, setIsModalOpen] = useState(false);
  //Modal.setAppElement("#root");
  console.log(isModalOpen);
  return(
    <div className="flex flex-col items-center border border-solid border-gray-200 py-4 transition-transform duration-200 hover:scale-110" onClick={() => setIsModalOpen(true)}>
      <img src={sketch.imageUrl} alt={sketch.imageUrl} className="w-1/2"/>
      <div>
        <div className='text-sm font-semibold text-gray-100'>
          {sketch.userName}
        </div>
        <div className='text-sm text-gray-200'>
          更新日時: {formatDateTime(sketch.updatedAt)}
        </div>
      </div>
      <div onClick={(event) => event.stopPropagation()}>
        <Modal isOpen={isModalOpen} onClose={() => setIsModalOpen(false)}>
          <img src={sketch.imageUrl} alt={sketch.imageUrl} />
          {/* <button onClick={() => console.log("clicked")}>閉じる</button> */}
        </Modal>
      </div>
    </div>
  );
}
