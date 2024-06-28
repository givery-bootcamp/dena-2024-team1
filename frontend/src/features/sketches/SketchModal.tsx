import { Modal } from "~/shared/components/Modal";
import { Sketch } from "~/generated";
type Props = {
  selectedSketch: Sketch;
  setSelectedSketch: (sketch: Sketch | null) => void;
};

export const SketchModal = ({ selectedSketch, setSelectedSketch }: Props) => {
  return (
    <Modal isOpen={true} onClose={() => setSelectedSketch(null)}>
      <img src={selectedSketch.image_url} alt={selectedSketch.image_url} />
      <h1>{selectedSketch.user_name}</h1>
    </Modal>
  );
};

