import { Modal } from "~/shared/components/Modal";
type Props = {
  selectedSketchUrl: string;
  setSelectedSketchUrl: (url: string | null) => void;
};

export const SketchModal = ({ selectedSketchUrl, setSelectedSketchUrl }: Props) => {
  return (
    <Modal isOpen={true} onClose={() => setSelectedSketchUrl(null)}>
      <img src={selectedSketchUrl} alt={selectedSketchUrl} />
    </Modal>
  );
};

