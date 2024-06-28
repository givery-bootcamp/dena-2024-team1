import { SketchColorPicker } from "./SketchColorPicker";

export const SketchTool = () => {
  return (
    <div className="flex w-56 flex-col gap-8">
      <h5 className="text-lg font-bold">お絵描きツール</h5>
      <div className="flex flex-col gap-6">
        <div className="flex flex-col items-center gap-2">
          <p className="w-full text-sm font-semibold text-gray-200">線の色</p>
          <SketchColorPicker />
        </div>
        <div className="flex flex-col items-center gap-2">
          <p className="w-full text-sm font-semibold text-gray-200">線の太さ</p>
          <p className="text-sm font-semibold">Coming soon...</p>
        </div>
      </div>
    </div>
  );
};
