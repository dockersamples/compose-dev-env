package handler

import (
	"aquafarm-management/app/model"
	"aquafarm-management/app/usecase"
	"github.com/gin-gonic/gin"
)

// ItemHandler menangani permintaan HTTP untuk item
type ItemHandler struct {
	itemUsecase *usecase.ItemUsecase
}

// NewItemHandler membuat instance baru ItemHandler
func NewItemHandler(iu *usecase.ItemUsecase) *ItemHandler {
	return &ItemHandler{
		itemUsecase: iu,
	}
}

// Fetch mengambil semua item
func (h *ItemHandler) Fetch(c *gin.Context) {
	// memanggil usecase untuk mengambil item
	items, err := h.itemUsecase.Fetch()
	if err != nil {
		// menangani error
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": items})
}

// Get mengambil item berdasarkan ID
func (h *ItemHandler) Get(c *gin.Context) {
	// mengambil ID item dari URL
	id := c.Param("id")
	// memanggil usecase untuk mengambil item
	item, err := h.itemUsecase.Get(id)
	if err != nil {
		// menangani error
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": item})
}

// Store membuat item baru
func (h *ItemHandler) Store(c *gin.Context) {
	// mengambil data item dari permintaan HTTP
	var item model.Item
	err := c.BindJSON(&item)
	if err != nil {
		// menangani error
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	// memanggil usecase untuk membuat item baru
	_, err = h.itemUsecase.Store(&item)
	if err != nil {
		// menangani error
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

// Update memperbarui item
func (h *ItemHandler) Update(c *gin.Context) {
	// mengambil ID item dari URL
	id := c.Param("id")
	// mengambil data item dari permintaan HTTP
	var item model.Item
	err := c.BindJSON(&item)
	if err != nil {
		// menangani error
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	// memanggil usecase untuk memperbarui item
	err = h.itemUsecase.Update(id, &item)
	if err != nil {
		// menangani error
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

// Delete menghapus item
func (h *ItemHandler) Delete(c *gin.Context) {
	// mengambil ID item dari URL
	id := c.Param("id")
	// memanggil usecase untuk menghapus item
	err := h.itemUsecase.Delete(id)
	if err != nil {
		// menangani error
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}
