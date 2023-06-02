package handlers

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/lctzz540/chessaiserver/models"
	"github.com/notnil/chess"
)

// fix bug here
func MakeMove(ctx context.Context, chessGame *chess.Game, idRoom string, move *models.Move, redisClient *redis.Client) error {
	err := chessGame.MoveStr(move.From + move.To)
	if err != nil {
		return err
	}

	if move.Promotion != "" {
		err := chessGame.MoveStr(move.From + move.To)
		if err != nil {
			return err
		}

		targetSquare, _ := chessGame.Position().Board().SquareByName(move.To)
		promotionPiece := chess.PieceFromFEN(move.Promotion)

		if targetSquare.Piece() != nil && targetSquare.Piece().Type() != promotionPiece.Type() {
			chessGame.Position().RemovePieceAt(targetSquare.Name())
			chessGame.Position().AddPiece(targetSquare.Name(), promotionPiece)
		}

		// Send move to Redis or handle it as desired
		_, err = redisClient.RPush(ctx, idRoom, move).Result()
		if err != nil {
			return err
		}

		return nil
	}
	return nil
}
func UpdateGameState(ctx context.Context, idRoom string, chessGame *chess.Game, redisClient *redis.Client) error {
	// Serialize the chess game state
	gameState := chessGame.String()

	// Set the game state in Redis using RPush command
	if err := redisClient.RPush(ctx, idRoom, gameState).Err(); err != nil {
		return err
	}

	return nil
}
func StartGame(ctx context.Context, roomId string, redisClient *redis.Client) *chess.Game {
	keyExists, err := redisClient.Exists(ctx, roomId).Result()
	if err != nil {
		return nil
	}

	if keyExists == 0 {
		err := redisClient.Set(ctx, roomId, "", 0).Err()
		if err != nil {
			return nil
		}
	} else {
		return nil
	}

	// Create a new chess game
	chessGame := chess.NewGame()

	err = UpdateGameState(ctx, roomId, chessGame, redisClient)
	if err != nil {
		return nil
	}

	return chessGame
}
