# クライアント用Makefile

# ローカルIPアドレスを自動取得
LOCAL_IP := $(shell route -n get default | grep interface | awk '{print $$2}' | xargs ifconfig | grep "inet " | head -1 | awk '{print $$2}')

.PHONY: env disable-google-services enable-google-services build-ios-dev build-android-dev build-ios-prd submit-ios help

# .env.localファイルを生成（ローカル開発用）
env:
	@echo "🔧 .env.localファイルを確認中..."
	@if [ -f .env.local ]; then \
		echo "📄 既存の.env.localファイルが見つかりました。"; \
		echo "📍 現在の設定:"; \
		cat .env.local; \
		echo ""; \
		echo "🔄 サーバー設定をローカル環境で上書き中..."; \
		sed -i '' 's/^EXPO_PUBLIC_SERVER_IP=.*/EXPO_PUBLIC_SERVER_IP=$(LOCAL_IP)/' .env.local; \
		echo "✅ サーバー設定が上書きされました:"; \
		cat .env.local; \
	else \
		echo "📄 新しい.env.localファイルを生成中..."; \
		echo "EXPO_PUBLIC_SERVER_IP=$(LOCAL_IP)" > .env.local; \
		echo "✅ .env.localファイルが生成されました:"; \
		cat .env.local; \
	fi

# .gitignoreのGoogleService-Info.plistをコメントアウト（ビルドに含める）
enable-google-services:
	@echo "🔧 GoogleService-Info.plistをビルドに含めるため.gitignoreを一時的に修正中..."
	@if grep -q "^\*\*/\*\.plist" .gitignore; then \
		sed -i '' 's/^\*\*\/\*\.plist/#**\/*.plist/' .gitignore; \
		echo "✅ .gitignoreの**/*.plistをコメントアウトしました"; \
	else \
		echo "✅ 既に**/*.plistはコメントアウト済みです"; \
	fi
	@if grep -q "^\*\*/google-services\.json" .gitignore; then \
		sed -i '' 's/^\*\*\/google-services\.json/#**\/google-services.json/' .gitignore; \
		echo "✅ .gitignoreの**/google-services.jsonをコメントアウトしました"; \
	else \
		echo "✅ 既に**/google-services.jsonはコメントアウト済みです"; \
	fi

# .gitignoreのGoogleService-Info.plistのコメントアウトを解除（元に戻す）
disable-google-services:
	@echo "🔧 .gitignoreを元に戻しています..."
	@if grep -q "^#\*\*/\*\.plist" .gitignore; then \
		sed -i '' 's/^#\*\*\/\*\.plist/**\/*.plist/' .gitignore; \
		echo "✅ .gitignoreの**/*.plistのコメントアウトを解除しました"; \
	else \
		echo "✅ 既に**/*.plistはコメントアウトされていません"; \
	fi
	@if grep -q "^#\*\*/google-services\.json" .gitignore; then \
		sed -i '' 's/^#\*\*\/google-services\.json/**\/google-services.json/' .gitignore; \
		echo "✅ .gitignoreの**/google-services.jsonのコメントアウトを解除しました"; \
	else \
		echo "✅ 既に**/google-services.jsonはコメントアウトされていません"; \
	fi

# EAS Buildコマンド（GoogleService-Info.plistを自動で含める）
build-ios-dev:
	@echo "🍎 iOS Development Buildを作成中..."
	@$(MAKE) enable-google-services
	@echo "📱 GoogleService-Info.plistをビルドに含めてビルド開始..."
	@eas build --platform ios --profile development-device || (echo "❌ ビルドに失敗しました。.gitignoreを元に戻します..." && $(MAKE) disable-google-services && exit 1)
	@$(MAKE) disable-google-services
	@echo "✅ ビルド完了！.gitignoreを元に戻しました"

build-android-dev:
	@echo "🤖 Android Development Buildを作成中..."
	@$(MAKE) enable-google-services
	@echo "📱 GoogleService-Info.plistをビルドに含めてビルド開始..."
	@eas build --platform android --profile development-device || (echo "❌ ビルドに失敗しました。.gitignoreを元に戻します..." && $(MAKE) disable-google-services && exit 1)
	@$(MAKE) disable-google-services
	@echo "✅ ビルド完了！.gitignoreを元に戻しました"

build-ios-preview:
	@echo "🍎 iOS Preview Buildを作成中..."
	@$(MAKE) enable-google-services
	@echo "📱 GoogleService-Info.plistをビルドに含めてビルド開始..."
	@eas build --platform ios --profile preview || (echo "❌ ビルドに失敗しました。.gitignoreを元に戻します..." && $(MAKE) disable-google-services && exit 1)
	@$(MAKE) disable-google-services
	@echo "✅ ビルド完了！.gitignoreを元に戻しました"

build-android-preview:
	@echo "🤖 Android Preview Buildを作成中..."
	@$(MAKE) enable-google-services
	@echo "📱 GoogleService-Info.plistをビルドに含めてビルド開始..."
	@eas build --platform android --profile preview || (echo "❌ ビルドに失敗しました。.gitignoreを元に戻します..." && $(MAKE) disable-google-services && exit 1)
	@$(MAKE) disable-google-services
	@echo "✅ ビルド完了！.gitignoreを元に戻しました"

build-ios-prd:
	@echo "🍎 iOS Production Buildを作成中..."
	@$(MAKE) enable-google-services
	@echo "📱 GoogleService-Info.plistをビルドに含めてビルド開始..."
	@eas build --platform ios --profile production || (echo "❌ ビルドに失敗しました。.gitignoreを元に戻します..." && $(MAKE) disable-google-services && exit 1)
	@$(MAKE) disable-google-services
	@echo "✅ ビルド完了！.gitignoreを元に戻しました"

submit-ios:
	@echo "🍎 iOS Production Buildを送信中..."
	eas submit --platform ios --profile production

# デフォルトターゲット
help:
	@echo "利用可能なコマンド:"
	@echo "  make env            - .env.localファイルを確認/生成（ローカルサーバー設定）"
	@echo ""
	@echo "EAS Build（クラウド上でビルド、GoogleService-Info.plistを自動で含める）:"
	@echo "  make build-ios-dev    - iOS Development Build作成"
	@echo "  make build-android-dev - Android Development Build作成"
	@echo "  make build-ios-preview - iOS Preview Build作成"
	@echo "  make build-android-preview - Android Preview Build作成"
	@echo "  make build-ios-prd    - iOS Production Build作成"
	@echo "  make submit-ios       - iOS Production Build送信"
	@echo ""
	@echo "手動操作:"
	@echo "  make enable-google-services  - .gitignoreの**/*.plistと**/google-services.jsonをコメントアウト"
	@echo "  make disable-google-services - .gitignoreの**/*.plistと**/google-services.jsonのコメントアウトを解除"
	@echo ""
	@echo "環境変数:"
	@echo "  LOCAL_IP        - 自動取得されたローカルIP: $(LOCAL_IP)"
