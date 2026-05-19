class Motivate < Formula
  desc "Print a random motivational quote with ASCII box framing"
  homepage "https://github.com/colbert677/motivational-quotes-cli"
  url "https://github.com/colbert677/motivational-quotes-cli/releases/download/v#{version}/motivate-macos-latest"
  license "MIT"
  version_scheme 1

  def install
    bin.install "motivate-macos-latest" => "motivate"
  end

  test do
    assert_match /—/, shell_output("#{bin}/motivate --no-color")
  end
end