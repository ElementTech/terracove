# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Terracove < Formula
  desc ""
  homepage "https://github.com/jatalocks/terracove"
  version "0.0.1"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/jatalocks/terracove/releases/download/v0.0.1/terracove_0.0.1_Darwin_arm64.tar.gz"
      sha256 "3a95466156338af0f3d8b330817bb21a646d50250d039784742213ba8da8a535"

      def install
        bin.install "terracove"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/jatalocks/terracove/releases/download/v0.0.1/terracove_0.0.1_Darwin_x86_64.tar.gz"
      sha256 "ce34dd9b96effab18aea4c66f7b0803c8f62de1328cb6ddd9fb2a55378679db6"

      def install
        bin.install "terracove"
      end
    end
  end

  on_linux do
    if Hardware::CPU.intel?
      url "https://github.com/jatalocks/terracove/releases/download/v0.0.1/terracove_0.0.1_Linux_x86_64.tar.gz"
      sha256 "f330a3677df0eea7709f0fe1b3912437e36405870998357c7ad522fba2d16ce2"

      def install
        bin.install "terracove"
      end
    end
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/jatalocks/terracove/releases/download/v0.0.1/terracove_0.0.1_Linux_arm64.tar.gz"
      sha256 "674e0f8709ce669af3753ebed05efe967c0dd33bba450b226ed2461186e952a5"

      def install
        bin.install "terracove"
      end
    end
  end
end