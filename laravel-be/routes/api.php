<?php

// routes/api.php

use App\Http\Controllers\SiswaApiController;

Route::apiResource('siswas', SiswaApiController::class);
