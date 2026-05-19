class Motivate < Formula
  desc "Display random motivational quotes in your terminal"
  homepage "https://github.com/$GITHUB_REPO/motivational-quotes-cli"
  url "https://github.com/$GITHUB_REPO/motivational-quotes-cli/archive/refs/tags/v1.0.0.tar.gz"
  sha256 "TODO: run 'brew formula .' after first tag to get sha256"
  license "MIT"
  version "1.0.0"

  def install
    system "go", "build", "-o", "motivate", "."
    bin.install "motivate"
  end

  test do
    assert_match /—/, shell_output("#{bin}/motivate --no-color")
  end
end
