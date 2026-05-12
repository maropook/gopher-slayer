# Gopher Slayer - Workshop Tasks

各レベルで「動かないゲームのバグを直す」ことで、GoのBackend開発の基礎を体験する。

---

## Lv1：ヒーローが攻撃しても0ダメージ

**症状：** 攻撃ボタンを押しても「You dealt 0 damage!」と表示され、敵のHPが減らない。

**空白箇所：** `internal/service/battle_service.go`

```go
// ダメージを計算する関数
// この関数を完成させてください
func calculateDamage(attack int) int {
    return 0 // ← ここを修正する
}
```

**やること：** `return 0` を、攻撃力をもとにダメージを返す処理に書き換える。

**完成イメージ：**
```go
func calculateDamage(attack int) int {
    // attackの値をそのまま返すだけでもOK！
    return attack
}
```

**体験できること：** 関数の役割・戻り値の理解

---

## Lv2：ステージをクリアしても経験値が増えない

**症状：** ステージをクリアすると「EXP +30」と画面には出るが、
リロードするとEXPが0のままになっている。

**空白箇所：** `internal/service/stage_service.go`

```go
func (s *StageService) ClearStage(stageID int) (*model.ClearStageResponse, error) {
    // ... 省略 ...

    newExp := hero.Experience + expGained

    // ← ここにDBへの保存処理が抜けている

    return &model.ClearStageResponse{
        Message:          fmt.Sprintf("Stage '%s' cleared!", stage.Name),
        ExperienceGained: expGained,
        NewExperience:    newExp,
    }, nil
}
```

**やること：** `heroRepo.UpdateExperience()` を呼び出す処理を追加する。
参考として、`internal/repository/hero_repository.go` の `UpdateName()` を見てみよう。

**完成イメージ：**
```go
// 4. DBに経験値を保存する
if err := s.heroRepo.UpdateExperience(newExp); err != nil {
    return nil, fmt.Errorf("failed to update experience: %w", err)
}
```

**体験できること：** DBへの書き込み（UPDATE）、Backendの醍醐味

---

## Lv3：ラストステージのボスが強すぎて詰んだ

**症状：** Stage5「Dragon's Lair」のBoss Dragonの攻撃力が50もあり、どうやっても勝てない。
ゲーム画面の「Edit HP」ボタンを押してHPを編集しようとすると、エラーになってしまう。
`PUT /api/hero/hp` というAPIを呼んでいるが、このエンドポイントが存在しないためだ。

**やること：** ヒーローのHPを編集できるAPIを、ルーティングから作成する。

---

### Step 1：ルートを登録する（`main.go`）

```go
// Hero routes
api.GET("/hero", heroHandler.GetHero)
api.PUT("/hero/name", heroHandler.UpdateName)
api.PUT("/hero/experience", heroHandler.UpdateExperience)
// ← ここにHP更新のルートを追加する
```

`heroHandler.UpdateHP` はすでに実装済み。以下の1行を追加しよう。

```go
api.PUT("/hero/hp", heroHandler.UpdateHP)
```

### Step 2：動作確認

ルートを追加したら、ゲーム画面の「Edit HP」ボタンでHPを編集してからボス戦に挑もう。

Swagger（`http://localhost:8080/docs/swagger.yaml`）や curl でも確認できる：

```bash
curl -X PUT http://localhost:8080/api/hero/hp \
  -H "Content-Type: application/json" \
  -d '{"hp": 100}'
```

### Step 3：コードの流れを追う（理解を深めたい人向け）

ルートを追加するだけでなく、リクエストがどう処理されるか流れを読んでみよう。

```
main.go（ルーティング）
  └─ internal/handler/hero_handler.go の UpdateHP()
       └─ internal/service/hero_service.go の UpdateHP()
            └─ internal/repository/hero_repository.go の UpdateHP()
                 └─ UPDATE heroes SET hp = ? WHERE id = 1
```

**体験できること：** ルーティング追加、handler → service → repository の全体の流れ

---

## Lv4：特定の敵の攻撃がおかしい

**症状：** Hell Gateステージの「Hell Hound」と戦うと、
攻撃が来るまで数秒かかり、しかもHPが増えてしまう（ダメージがマイナスになっている）。

**空白箇所：** `internal/service/battle_service.go`

```go
func (s *BattleService) EnemyAttack(req EnemyAttackRequest) AttackResponse {
    // ← バグが仕込まれている。コードをよく読んで見つけよう。
    damage := calculateDamage(req.EnemyAttack)
    return AttackResponse{
        Damage:  damage,
        Message: fmt.Sprintf("%s dealt %d damage!", req.EnemyName, damage),
    }
}
```

**やること：** バグを自分で見つけて修正する。
- ヒント1：なぜ攻撃が遅いのか？
- ヒント2：なぜHPが増えてしまうのか？

**体験できること：** デバッグ力、処理の流れを追う読解力

---

## Lv5（応用）：ボスのステータスを変更するAPIがない

**症状：** Boss Dragonのステータスが弱すぎる（または強すぎる）。
Swaggerを見ると `PUT /api/enemies/:id` の仕様が書かれているが、実装されていない。

**やること：** ルーティング・handler・service・repositoryをゼロから追加する。

**Swagger仕様：**
```
PUT /api/enemies/:id
Request Body:
  {
    "hp":     300,
    "max_hp": 300,
    "attack": 50
  }
Response:
  { "message": "Enemy updated successfully" }
```

**実装すること：**
1. `internal/repository/stage_repository.go` に `UpdateEnemy()` を追加
2. `internal/service/stage_service.go` に `UpdateEnemy()` を追加
3. `internal/handler/stage_handler.go` に `UpdateEnemy()` ハンドラーを追加
4. `main.go` にルートを追加: `api.PUT("/enemies/:id", stageHandler.UpdateEnemy)`

**バリデーション（余力があれば）：** HPや攻撃力が0以下のリクエストはエラーを返す。

**体験できること：** APIをゼロから設計・実装する全体像、バリデーション

---

## 参考：実装済みの例

迷ったときは以下の既存コードを参考にしよう。

| 参考にできる実装 | ファイル |
|----------------|---------|
| DB更新の書き方（UPDATE） | `internal/repository/hero_repository.go` の `UpdateName()` |
| ハンドラーの書き方 | `internal/handler/hero_handler.go` の `UpdateName()` |
| サービス層の書き方 | `internal/service/hero_service.go` の `UpdateName()` |
| ルーティングの追加 | `main.go` の `api.PUT("/hero/name", ...)` |
