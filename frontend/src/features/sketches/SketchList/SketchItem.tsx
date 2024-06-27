import { useState }from "react";

import { Modal } from "~/shared/components/Modal";
import type { Sketch } from "~/generated";

type SketchItemProps = {
  sketch: Sketch
};

export const SketchItem = ({ sketch }: SketchItemProps) => {
  const [isModalOpen, setIsModalOpen] = useState(false);
  return (
    <div onClick={() => setIsModalOpen(true)}>
      <img className="h-[200px] w-[300px] border border-gray-200 object-cover" src={sketch.imageUrl} alt={sketch.imageUrl} />
      <div onClick={(event) => event.stopPropagation()}>
        <Modal isOpen={isModalOpen} onClose={() => setIsModalOpen(false)}>
          <img src={sketch.imageUrl} alt={sketch.imageUrl} />
          {/* <button onClick={() => console.log("clicked")}>閉じる</button> */}
        </Modal>
      </div>
    </div>
  );
};
