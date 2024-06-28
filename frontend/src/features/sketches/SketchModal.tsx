import { Modal } from "~/shared/components/Modal";
import { Sketch } from "~/generated";

type Props = {
  selectedSketch: Sketch;
  setSelectedSketch: (sketch: Sketch | null) => void;
};

export const SketchModal = ({ selectedSketch, setSelectedSketch }: Props) => {
  return (
    <Modal isOpen={true} onClose={() => setSelectedSketch(null)}>
      <div className="flex h-full flex-col">
        <img src={selectedSketch.image_url} alt={selectedSketch.image_url} />
        <div className="absolute bottom-2 right-2 text-xl font-semibold text-gray-200">
          {selectedSketch.user_name}
        </div>
      </div>
    </Modal>
  );
};
